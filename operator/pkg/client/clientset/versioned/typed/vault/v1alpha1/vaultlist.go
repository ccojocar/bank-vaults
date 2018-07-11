// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/banzaicloud/bank-vaults/operator/pkg/apis/vault/v1alpha1"
	scheme "github.com/banzaicloud/bank-vaults/operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VaultListsGetter has a method to return a VaultListInterface.
// A group's client should implement this interface.
type VaultListsGetter interface {
	VaultLists(namespace string) VaultListInterface
}

// VaultListInterface has methods to work with VaultList resources.
type VaultListInterface interface {
	Create(*v1alpha1.VaultList) (*v1alpha1.VaultList, error)
	Update(*v1alpha1.VaultList) (*v1alpha1.VaultList, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VaultList, error)
	List(opts v1.ListOptions) (*v1alpha1.VaultListList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VaultList, err error)
	VaultListExpansion
}

// vaultLists implements VaultListInterface
type vaultLists struct {
	client rest.Interface
	ns     string
}

// newVaultLists returns a VaultLists
func newVaultLists(c *VaultV1alpha1Client, namespace string) *vaultLists {
	return &vaultLists{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vaultList, and returns the corresponding vaultList object, and an error if there is any.
func (c *vaultLists) Get(name string, options v1.GetOptions) (result *v1alpha1.VaultList, err error) {
	result = &v1alpha1.VaultList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vaultlists").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VaultLists that match those selectors.
func (c *vaultLists) List(opts v1.ListOptions) (result *v1alpha1.VaultListList, err error) {
	result = &v1alpha1.VaultListList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vaultlists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vaultLists.
func (c *vaultLists) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vaultlists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a vaultList and creates it.  Returns the server's representation of the vaultList, and an error, if there is any.
func (c *vaultLists) Create(vaultList *v1alpha1.VaultList) (result *v1alpha1.VaultList, err error) {
	result = &v1alpha1.VaultList{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vaultlists").
		Body(vaultList).
		Do().
		Into(result)
	return
}

// Update takes the representation of a vaultList and updates it. Returns the server's representation of the vaultList, and an error, if there is any.
func (c *vaultLists) Update(vaultList *v1alpha1.VaultList) (result *v1alpha1.VaultList, err error) {
	result = &v1alpha1.VaultList{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vaultlists").
		Name(vaultList.Name).
		Body(vaultList).
		Do().
		Into(result)
	return
}

// Delete takes name of the vaultList and deletes it. Returns an error if one occurs.
func (c *vaultLists) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vaultlists").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vaultLists) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vaultlists").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched vaultList.
func (c *vaultLists) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VaultList, err error) {
	result = &v1alpha1.VaultList{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vaultlists").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}