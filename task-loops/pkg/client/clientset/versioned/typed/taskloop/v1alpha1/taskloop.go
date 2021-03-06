/*
Copyright 2020 The Tekton Authors

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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/tektoncd/experimental/task-loops/pkg/apis/taskloop/v1alpha1"
	scheme "github.com/tektoncd/experimental/task-loops/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TaskLoopsGetter has a method to return a TaskLoopInterface.
// A group's client should implement this interface.
type TaskLoopsGetter interface {
	TaskLoops(namespace string) TaskLoopInterface
}

// TaskLoopInterface has methods to work with TaskLoop resources.
type TaskLoopInterface interface {
	Create(*v1alpha1.TaskLoop) (*v1alpha1.TaskLoop, error)
	Update(*v1alpha1.TaskLoop) (*v1alpha1.TaskLoop, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.TaskLoop, error)
	List(opts v1.ListOptions) (*v1alpha1.TaskLoopList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TaskLoop, err error)
	TaskLoopExpansion
}

// taskLoops implements TaskLoopInterface
type taskLoops struct {
	client rest.Interface
	ns     string
}

// newTaskLoops returns a TaskLoops
func newTaskLoops(c *CustomV1alpha1Client, namespace string) *taskLoops {
	return &taskLoops{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the taskLoop, and returns the corresponding taskLoop object, and an error if there is any.
func (c *taskLoops) Get(name string, options v1.GetOptions) (result *v1alpha1.TaskLoop, err error) {
	result = &v1alpha1.TaskLoop{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("taskloops").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TaskLoops that match those selectors.
func (c *taskLoops) List(opts v1.ListOptions) (result *v1alpha1.TaskLoopList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TaskLoopList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("taskloops").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested taskLoops.
func (c *taskLoops) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("taskloops").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a taskLoop and creates it.  Returns the server's representation of the taskLoop, and an error, if there is any.
func (c *taskLoops) Create(taskLoop *v1alpha1.TaskLoop) (result *v1alpha1.TaskLoop, err error) {
	result = &v1alpha1.TaskLoop{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("taskloops").
		Body(taskLoop).
		Do().
		Into(result)
	return
}

// Update takes the representation of a taskLoop and updates it. Returns the server's representation of the taskLoop, and an error, if there is any.
func (c *taskLoops) Update(taskLoop *v1alpha1.TaskLoop) (result *v1alpha1.TaskLoop, err error) {
	result = &v1alpha1.TaskLoop{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("taskloops").
		Name(taskLoop.Name).
		Body(taskLoop).
		Do().
		Into(result)
	return
}

// Delete takes name of the taskLoop and deletes it. Returns an error if one occurs.
func (c *taskLoops) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("taskloops").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *taskLoops) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("taskloops").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched taskLoop.
func (c *taskLoops) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TaskLoop, err error) {
	result = &v1alpha1.TaskLoop{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("taskloops").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
