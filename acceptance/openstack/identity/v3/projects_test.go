// +build acceptance

package v3

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
)

func TestProjectsList(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	var iTrue bool = true
	listOpts := projects.ListOpts{
		Enabled: &iTrue,
	}

	allPages, err := projects.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to list projects: %v", err)
	}

	allProjects, err := projects.ExtractProjects(allPages)
	if err != nil {
		t.Fatalf("Unable to extract projects: %v", err)
	}

	for _, project := range allProjects {
		PrintProject(t, &project)
	}
}

func TestProjectsGet(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	allPages, err := projects.List(client, nil).AllPages()
	if err != nil {
		t.Fatalf("Unable to list projects: %v", err)
	}

	allProjects, err := projects.ExtractProjects(allPages)
	if err != nil {
		t.Fatalf("Unable to extract projects: %v", err)
	}

	project := allProjects[0]
	p, err := projects.Get(client, project.ID, nil).Extract()
	if err != nil {
		t.Fatalf("Unable to get project: %v", err)
	}

	PrintProject(t, p)
}

func TestProjectsCRUD(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v")
	}

	project, err := CreateProject(t, client, nil)
	if err != nil {
		t.Fatalf("Unable to create project: %v", err)
	}
	defer DeleteProject(t, client, project.ID)

	PrintProject(t, project)

	var iFalse bool = false
	updateOpts := projects.UpdateOpts{
		Enabled: &iFalse,
	}

	updatedProject, err := projects.Update(client, project.ID, updateOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to update project: %v", err)
	}

	PrintProject(t, updatedProject)
}

func TestProjectsDomain(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v")
	}

	var iTrue = true
	createOpts := projects.CreateOpts{
		IsDomain: &iTrue,
	}

	projectDomain, err := CreateProject(t, client, &createOpts)
	if err != nil {
		t.Fatalf("Unable to create project: %v", err)
	}
	defer DeleteProject(t, client, projectDomain.ID)

	PrintProject(t, projectDomain)

	createOpts = projects.CreateOpts{
		DomainID: projectDomain.ID,
	}

	project, err := CreateProject(t, client, &createOpts)
	if err != nil {
		t.Fatalf("Unable to create project: %v", err)
	}
	defer DeleteProject(t, client, project.ID)

	PrintProject(t, project)

	var iFalse = false
	updateOpts := projects.UpdateOpts{
		Enabled: &iFalse,
	}

	_, err = projects.Update(client, projectDomain.ID, updateOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to disable domain: %v")
	}
}

func TestProjectsNested(t *testing.T) {
	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v")
	}

	projectMain, err := CreateProject(t, client, nil)
	if err != nil {
		t.Fatalf("Unable to create project: %v", err)
	}
	defer DeleteProject(t, client, projectMain.ID)

	PrintProject(t, projectMain)

	createOpts := projects.CreateOpts{
		ParentID: projectMain.ID,
	}

	project, err := CreateProject(t, client, &createOpts)
	if err != nil {
		t.Fatalf("Unable to create project: %v", err)
	}
	defer DeleteProject(t, client, project.ID)

	PrintProject(t, project)
}
