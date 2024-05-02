import hglib
from messages import RunnerStatusMessage


class MercurialBackportRunnerPipeline:
    def __init__(self, source_path, target_branch_name, new_branch_name, commits, reporter) -> None:
        self.target_branch_name = target_branch_name
        self.new_branch_name = new_branch_name
        self.commits = commits

        self.hg = hglib.open(source_path)
        self.reporter = reporter

    def run(self):
        self.pull()
        self.checkout(self.target_branch_name, clean=True)
        self.checkout_new_branch(self.new_branch_name)
        for commit in self.commits:
            self.cherry_pick(commit)
        self.push(self.new_branch_name)

    def checkout(self, branch, clean=False):
        print('Checking out branch: ' + branch)
        self.reporter.send(RunnerStatusMessage(
            'checkout', 'start', {"branch": branch}).get())
        try:
            self.git.checkout(branch, force=True, clean=clean)
            self.reporter.send(RunnerStatusMessage(
                'checkout', 'success', {"branch": branch}).get())
        except Exception as e:
            print(e)
            self.reporter.send(RunnerStatusMessage(
                'checkout', 'failure', {"branch": branch, "error": str(e)}).get())
            exit(1)

    def checkout_new_branch(self, branch):
        print('Checking out new branch: ' + branch)
        self.reporter.send(RunnerStatusMessage(
            'checkout_new_branch', 'start', {"branch": branch}).get())
        try:
            self.hg.branch(branch)
            self.reporter.send(RunnerStatusMessage(
                'checkout_new_branch', 'success', {"branch": branch}).get())
        except Exception as e:
            print(e)
            self.reporter.send(RunnerStatusMessage(
                'checkout_new_branch', 'failure', {"branch": branch, "error": str(e)}).get())
            exit(1)

    def cherry_pick(self, revision):
        print('Cherry picking commit: ' + revision)
        self.reporter.send(RunnerStatusMessage(
            'cherry_pick', 'start', {"commit": revision}).get())
        try:
            self.hg.graft(revision)
            self.reporter.send(RunnerStatusMessage(
                'cherry_pick', 'success', {"commit": revision}).get())
        except Exception as e:
            print(f'Failed to cherry pick commit {revision}')
            print(e)
            self.reporter.send(RunnerStatusMessage('cherry_pick', 'failure', {
                "commit": revision, "error": str(e)}).get())
            exit(1)

    def pull(self):
        print('Pulling from origin')
        self.reporter.send(RunnerStatusMessage(
            'pull', 'start', {}).get()
        )
        try:
            self.hg.pull()
            self.reporter.send(RunnerStatusMessage(
                'pull', 'success', {}).get()
            )
        except Exception as e:
            print('Failed to pull from origin')
            print(e)
            self.reporter.send(RunnerStatusMessage(
                'pull', 'failure', {"error": str(e)}).get()
            )
            exit(1)

    def push(self, branch):
        print('Pushing to origin')
        self.reporter.send(RunnerStatusMessage(
            'push', 'start', {"branch": branch}).get()
        )
        try:
            self.hg.push(branch=branch, newbranch=True)
            self.reporter.send(RunnerStatusMessage(
                'push', 'success', {"branch": branch}).get()
            )
        except Exception as e:
            print('Failed to push to origin')
            self.reporter.send(RunnerStatusMessage(
                'push', 'failure', {"branch": branch, "error": str(e)}).get()
            )
            exit(1)


class MergeConflictException(Exception):
    pass
