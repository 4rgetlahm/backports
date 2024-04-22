import git
from messages import RunnerStatusMessage


class BackportRunnerPipeline:
    def __init__(self, source_path, target_branch_name, new_branch_name, commits, reporter) -> None:
        self.target_branch_name = target_branch_name
        self.new_branch_name = new_branch_name
        self.commits = commits

        self.git = git.Git(source_path)
        self.reporter = reporter
        self.push_retry_count = 0
        self.push_retry_count_limit = 5

        repo = git.Repo(source_path)

        repo.config_writer().set_value("user", "name", "Backport Automation").release()
        repo.config_writer().set_value("user", "email", "backport-automation@backports.com").release()

    def run(self):
        self.fetch()
        self.checkout(self.target_branch_name, clean=True)
        self.pull()
        self.checkout_new_branch(self.new_branch_name)
        for commit in self.commits:
            self.cherry_pick(commit, ['-m 1'])
        self.push()

    def checkout(self, branch, clean=False):
        print('Checking out branch: ' + branch)
        self.reporter.send(RunnerStatusMessage(
            'checkout', 'start', {"branch": branch}).get())
        try:
            if clean:
                self.git.reset('--hard')

            self.git.checkout(branch)
            
            if clean:
                self.git.reset('--hard')
                self.git.clean('-dfx')
            self.reporter.send(RunnerStatusMessage(
                'checkout', 'success', {"branch": branch}).get())
        except git.GitCommandError as e:
            print(
                f'Failed to check out to branch {branch}. Please make sure that branch exists.')
            print(e)
            self.reporter.send(RunnerStatusMessage(
                'checkout', 'failure', {"branch": branch, "error": str(e)}).get())
            exit(1)

    def checkout_new_branch(self, branch):
        print('Checking out new branch: ' + branch)
        self.reporter.send(RunnerStatusMessage(
            'checkout_new_branch', 'start', {"branch": branch}).get())
        try:
            self.git.checkout('-b', branch)
            self.reporter.send(RunnerStatusMessage(
                'checkout_new_branch', 'success', {"branch": branch}).get())
        except git.GitCommandError as e:
            print(
                f'Failed to check out to new branch {branch}. Please make sure that branch does not exist.')
            print
            self.reporter.send(RunnerStatusMessage(
                'checkout_new_branch', 'failure', {"branch": branch, "error": str(e)}).get())
            exit(1)

    def cherry_pick(self, commit, options):
        print('Cherry picking commit: ' + commit)
        self.reporter.send(RunnerStatusMessage(
            'cherry_pick', 'start', {"commit": commit}).get())
        try:
            self.git.cherry_pick(options, commit)
            self.reporter.send(RunnerStatusMessage(
                'cherry_pick', 'success', {"commit": commit}).get())
        except git.GitCommandError as e:
            if "CONFLICT" in e.stdout:
                print('Merge conflict was detected')
                print(e)
                self.reporter.send(RunnerStatusMessage('cherry_pick', 'failure', {
                    "commit": commit, "error": str(e)}).get())
                raise MergeConflictException
            else:
                print(f'Failed to cherry pick commit {commit}')
                print(e)
                self.reporter.send(RunnerStatusMessage('cherry_pick', 'failure', {
                    "commit": commit, "error": str(e)}).get())
                exit(1)

    def pull(self):
        print('Pulling from origin')
        self.reporter.send(RunnerStatusMessage(
            'pull', 'start', {}).get()
        )
        try:
            self.git.pull()
            self.reporter.send(RunnerStatusMessage(
                'pull', 'success', {}).get()
            )
        except git.GitCommandError as e:
            print('Failed to pull from origin')
            print(e)
            self.reporter.send(RunnerStatusMessage(
                'pull', 'failure', {"error": str(e)}).get()
            )
            exit(1)

    def fetch(self):
        print('Fetching from origin')
        self.reporter.send(RunnerStatusMessage(
            'fetch', 'start', {}).get()
        )
        try:
            self.git.fetch('--prune', '--all')
            self.reporter.send(RunnerStatusMessage(
                'fetch', 'success', {}).get()
            )
        except git.GitCommandError as e:
            print('Failed to fetch from origin')
            print(e)
            self.reporter.send(RunnerStatusMessage(
                'fetch', 'failure', {"error": str(e)}).get()
            )
            exit(1)

    def push(self):
        print('Pushing to origin')
        self.reporter.send(RunnerStatusMessage(
            'push', 'start', {"branch": self.new_branch_name}).get()
        )
        try:
            self.git.push('--set-upstream', 'origin', self.new_branch_name)
            self.reporter.send(RunnerStatusMessage(
                'push', 'success', {"branch": self.new_branch_name}).get()
            )
        except git.GitCommandError as e:
            print('Failed to push to origin')
            print(e)
            if self.push_retry_count < self.push_retry_count_limit:
                self.push_retry_count += 1
                self.push()
                return
            self.reporter.send(RunnerStatusMessage(
                'push', 'failure', {"branch": self.new_branch_name, "error": str(e)}).get()
            )
            exit(1)


class MergeConflictException(Exception):
    pass
