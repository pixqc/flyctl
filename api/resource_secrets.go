package api

func (c *Client) SetSecrets(appName string, secrets map[string]string) (*Release, error) {
	query := `
		mutation($input: SetSecretsInput!) {
			setSecrets(input: $input) {
				release {
					id
					version
					reason
					description
					user {
						id
						email
						name
					}
					createdAt
				}
			}
		}
	`

	input := SetSecretsInput{AppID: appName}
	for k, v := range secrets {
		input.Secrets = append(input.Secrets, SetSecretsInputSecret{Key: k, Value: v})
	}

	req := c.NewRequest(query)

	req.Var("input", input)

	data, err := c.Run(req)
	if err != nil {
		return nil, err
	}

	return &data.SetSecrets.Release, nil
}

func (c *Client) UnsetSecrets(appName string, keys []string) (*Release, error) {
	query := `
		mutation($input: UnsetSecretsInput!) {
			unsetSecrets(input: $input) {
				release {
					id
					version
					reason
					description
					user {
						id
						email
						name
					}
					createdAt
				}
			}
		}
	`

	req := c.NewRequest(query)

	req.Var("input", UnsetSecretsInput{AppID: appName, Keys: keys})

	data, err := c.Run(req)
	if err != nil {
		return nil, err
	}

	return &data.UnsetSecrets.Release, nil
}

func (c *Client) GetAppSecrets(appName string) ([]Secret, error) {
	query := `
		query ($appName: String!) {
			app(name: $appName) {
				secrets {
					name
					digest
					createdAt
				}
			}
		}
	`

	req := c.NewRequest(query)

	req.Var("appName", appName)

	data, err := c.Run(req)
	if err != nil {
		return nil, err
	}

	return data.App.Secrets, nil
}
