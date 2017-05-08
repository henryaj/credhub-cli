package commands

import (
	"github.com/cloudfoundry-incubator/credhub-cli/actions"
	"github.com/cloudfoundry-incubator/credhub-cli/client"
	"github.com/cloudfoundry-incubator/credhub-cli/config"
	"github.com/cloudfoundry-incubator/credhub-cli/models"
	"github.com/cloudfoundry-incubator/credhub-cli/repositories"
)

type GetCommand struct {
	CredentialIdentifier string `short:"n" required:"yes" long:"name" description:"Name of the credential to retrieve"`
	OutputJson           bool   `long:"output-json" description:"Return response in JSON format"`
}

func (cmd GetCommand) Execute([]string) error {
	cfg := config.ReadConfig()
	repository := repositories.NewCredentialRepository(client.NewHttpClient(cfg))
	action := actions.NewAction(repository, cfg)
	credential, err := action.DoAction(client.NewGetCredentialRequest(cfg, cmd.CredentialIdentifier), cmd.CredentialIdentifier)
	if err != nil {
		return err
	}

	models.Println(credential, cmd.OutputJson)

	return nil
}
