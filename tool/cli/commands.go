// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "env": CLI Commands
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-env/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-env-client
// --pkg=env
// --version=v1.3.0

package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fabric8-services/fabric8-env-client/env"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// CreateEnvironmentCommand is the command line data structure for the create action of environment
	CreateEnvironmentCommand struct {
		Payload     string
		ContentType string
		// ID of the space
		SpaceID     string
		PrettyPrint bool
	}

	// ListEnvironmentCommand is the command line data structure for the list action of environment
	ListEnvironmentCommand struct {
		// ID of the space
		SpaceID     string
		PrettyPrint bool
	}

	// ShowEnvironmentCommand is the command line data structure for the show action of environment
	ShowEnvironmentCommand struct {
		// ID of the environment
		EnvID       string
		PrettyPrint bool
	}

	// ShowStatusCommand is the command line data structure for the show action of status
	ShowStatusCommand struct {
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *env.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: `Create environment`,
	}
	tmp1 := new(CreateEnvironmentCommand)
	sub = &cobra.Command{
		Use:   `environment ["/api/spaces/SPACEID/environments"]`,
		Short: ``,
		Long: `

Payload example:

{
   "data": {
      "attributes": {
         "cluster-url": "https://api.starter-us-east-2a.openshift.com",
         "name": "myapp-stage",
         "namespaceName": "myapp-stage",
         "type": "stage"
      },
      "id": "40bbdd3d-8b5d-4fd6-ac90-7236b669af04",
      "links": {
         "meta": {
            "Ipsam non nesciunt.": false
         },
         "related": "Fugiat dolorem nostrum voluptas libero.",
         "self": "Sint optio et adipisci vero fugiat."
      },
      "type": "environments"
   },
   "included": [
      "dfee74cb-9986-4135-a70d-4688967f57c3"
   ]
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `List environments for the given space ID.`,
	}
	tmp2 := new(ListEnvironmentCommand)
	sub = &cobra.Command{
		Use:   `environment ["/api/spaces/SPACEID/environments"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `show action`,
	}
	tmp3 := new(ShowEnvironmentCommand)
	sub = &cobra.Command{
		Use:   `environment ["/api/environments/ENVID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp4 := new(ShowStatusCommand)
	sub = &cobra.Command{
		Use:   `status ["/api/status"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run makes the HTTP request corresponding to the CreateEnvironmentCommand command.
func (cmd *CreateEnvironmentCommand) Run(c *env.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/spaces/%v/environments", cmd.SpaceID)
	}
	var payload env.CreateEnvironmentPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateEnvironment(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateEnvironmentCommand) RegisterFlags(cc *cobra.Command, c *env.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var spaceID string
	cc.Flags().StringVar(&cmd.SpaceID, "spaceID", spaceID, `ID of the space`)
}

// Run makes the HTTP request corresponding to the ListEnvironmentCommand command.
func (cmd *ListEnvironmentCommand) Run(c *env.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/spaces/%v/environments", cmd.SpaceID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListEnvironment(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListEnvironmentCommand) RegisterFlags(cc *cobra.Command, c *env.Client) {
	var spaceID string
	cc.Flags().StringVar(&cmd.SpaceID, "spaceID", spaceID, `ID of the space`)
}

// Run makes the HTTP request corresponding to the ShowEnvironmentCommand command.
func (cmd *ShowEnvironmentCommand) Run(c *env.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/environments/%v", cmd.EnvID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowEnvironment(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowEnvironmentCommand) RegisterFlags(cc *cobra.Command, c *env.Client) {
	var envID string
	cc.Flags().StringVar(&cmd.EnvID, "envID", envID, `ID of the environment`)
}

// Run makes the HTTP request corresponding to the ShowStatusCommand command.
func (cmd *ShowStatusCommand) Run(c *env.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/status"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowStatus(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowStatusCommand) RegisterFlags(cc *cobra.Command, c *env.Client) {
}
