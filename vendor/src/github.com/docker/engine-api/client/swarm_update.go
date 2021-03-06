package client

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/docker/engine-api/types/swarm"
	"golang.org/x/net/context"
)

// SwarmUpdate updates the Swarm.
func (cli *Client) SwarmUpdate(ctx context.Context, version swarm.Version, swarm swarm.Spec, flags swarm.UpdateFlags) error {
	query := url.Values{}
	query.Set("version", strconv.FormatUint(version.Index, 10))
	query.Set("rotate_worker_token", fmt.Sprintf("%v", flags.RotateWorkerToken))
	query.Set("rotate_manager_token", fmt.Sprintf("%v", flags.RotateManagerToken))
	resp, err := cli.post(ctx, "/swarm/update", query, swarm, nil)
	ensureReaderClosed(resp)
	return err
}
