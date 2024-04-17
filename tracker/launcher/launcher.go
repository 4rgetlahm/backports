package launcher

import (
	"context"
	"fmt"
	"strings"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func LaunchBackportJob(clientset *kubernetes.Clientset, image string, reference string, newBranchName string, targetBranchName string, commits []string) string {
	jobClient := clientset.BatchV1().Jobs("default")
	var backoffLimit int32 = 0

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "backport-job-" + reference,
		},
		Spec: batchv1.JobSpec{
			BackoffLimit: &backoffLimit,
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "backport-runner",
							Image: image,
							Env: []corev1.EnvVar{
								{
									Name:  "REFERENCE",
									Value: reference,
								},
								{
									Name:  "NEW_BRANCH_NAME",
									Value: newBranchName,
								},
								{
									Name:  "TARGET_BRANCH_NAME",
									Value: targetBranchName,
								},
								{
									Name:  "COMMITS",
									Value: strings.Join(commits, ","),
								},
								{
									Name:  "SOURCE_PATH",
									Value: "/repo",
								},
								{
									Name:  "REPORTER_CONFIG",
									Value: "",
								},
							},
						},
					},
					RestartPolicy: corev1.RestartPolicyNever,
				},
			},
		},
	}

	fmt.Println("Creating job for backport (", reference, ")")

	result, err := jobClient.Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Job created for backport (", reference, ")")
	return result.GetObjectMeta().GetName()
}
