package gitlab

import (
	"fmt"
	"net/http"
	"testing"
)

func TestReleaseLinksService_ListReleaseLinks(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1/assets/links",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, exampleReleaseLinkList)
		})

	releaseLinks, _, err := client.ReleaseLinks.ListReleaseLinks(
		1, exampleTagName, &ListReleaseLinksOptions{},
	)
	if err != nil {
		t.Error(err)
	}
	if len(releaseLinks) != 2 {
		t.Error("expected 2 links")
	}
	if releaseLinks[0].Name != "awesome-v0.2.msi" {
		t.Errorf("release link name, expected '%s', got '%s'", "awesome-v0.2.msi",
			releaseLinks[0].Name)
	}
}

func TestReleaseLinksService_CreateReleaseLink(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1/assets/links",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			fmt.Fprint(w, exampleReleaseLink)
		})

	releaseLink, _, err := client.ReleaseLinks.CreateReleaseLink(
		1, exampleTagName,
		&CreateReleaseLinkOptions{
			Name: String(exampleReleaseName),
			URL:  String("http://192.168.10.15:3000"),
		})
	if err != nil {
		t.Error(err)
	}
	if releaseLink.Name != exampleReleaseName {
		t.Errorf("release link name, expected '%s', got '%s'", exampleReleaseName,
			releaseLink.Name)
	}
}

func TestReleaseLinksService_GetReleaseLink(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1/assets/links/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, exampleReleaseLink)
		})

	releaseLink, _, err := client.ReleaseLinks.GetReleaseLink(1, exampleTagName, 1)
	if err != nil {
		t.Error(err)
	}
	if releaseLink.Name != exampleReleaseName {
		t.Errorf("release link name, expected '%s', got '%s'", exampleReleaseName,
			releaseLink.Name)
	}
}

func TestReleaseLinksService_UpdateReleaseLink(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1/assets/links/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "PUT")
			fmt.Fprint(w, exampleReleaseLink)
		})

	releaseLink, _, err := client.ReleaseLinks.UpdateReleaseLink(
		1, exampleTagName, 1,
		&UpdateReleaseLinkOptions{
			Name: String(exampleReleaseName),
		})
	if err != nil {
		t.Error(err)
	}
	if releaseLink.Name != exampleReleaseName {
		t.Errorf("release link name, expected '%s', got '%s'", exampleReleaseName,
			releaseLink.Name)
	}
}

func TestReleaseLinksService_DeleteReleaseLink(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/projects/1/releases/v0.1/assets/links/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
			fmt.Fprint(w, exampleReleaseLink)
		})

	releaseLink, _, err := client.ReleaseLinks.DeleteReleaseLink(1, exampleTagName, 1)
	if err != nil {
		t.Error(err)
	}
	if releaseLink.Name != exampleReleaseName {
		t.Errorf("release link name, expected '%s', got '%s'", exampleReleaseName,
			releaseLink.Name)
	}
}
