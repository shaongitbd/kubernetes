/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha3

import (
	"context"

	v1alpha3 "k8s.io/api/resource/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	resourcev1alpha3 "k8s.io/client-go/applyconfigurations/resource/v1alpha3"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// PodSchedulingContextsGetter has a method to return a PodSchedulingContextInterface.
// A group's client should implement this interface.
type PodSchedulingContextsGetter interface {
	PodSchedulingContexts(namespace string) PodSchedulingContextInterface
}

// PodSchedulingContextInterface has methods to work with PodSchedulingContext resources.
type PodSchedulingContextInterface interface {
	Create(ctx context.Context, podSchedulingContext *v1alpha3.PodSchedulingContext, opts v1.CreateOptions) (*v1alpha3.PodSchedulingContext, error)
	Update(ctx context.Context, podSchedulingContext *v1alpha3.PodSchedulingContext, opts v1.UpdateOptions) (*v1alpha3.PodSchedulingContext, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, podSchedulingContext *v1alpha3.PodSchedulingContext, opts v1.UpdateOptions) (*v1alpha3.PodSchedulingContext, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha3.PodSchedulingContext, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha3.PodSchedulingContextList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.PodSchedulingContext, err error)
	Apply(ctx context.Context, podSchedulingContext *resourcev1alpha3.PodSchedulingContextApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha3.PodSchedulingContext, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, podSchedulingContext *resourcev1alpha3.PodSchedulingContextApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha3.PodSchedulingContext, err error)
	PodSchedulingContextExpansion
}

// podSchedulingContexts implements PodSchedulingContextInterface
type podSchedulingContexts struct {
	*gentype.ClientWithListAndApply[*v1alpha3.PodSchedulingContext, *v1alpha3.PodSchedulingContextList, *resourcev1alpha3.PodSchedulingContextApplyConfiguration]
}

// newPodSchedulingContexts returns a PodSchedulingContexts
func newPodSchedulingContexts(c *ResourceV1alpha3Client, namespace string) *podSchedulingContexts {
	return &podSchedulingContexts{
		gentype.NewClientWithListAndApply[*v1alpha3.PodSchedulingContext, *v1alpha3.PodSchedulingContextList, *resourcev1alpha3.PodSchedulingContextApplyConfiguration](
			"podschedulingcontexts",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha3.PodSchedulingContext { return &v1alpha3.PodSchedulingContext{} },
			func() *v1alpha3.PodSchedulingContextList { return &v1alpha3.PodSchedulingContextList{} }),
	}
}