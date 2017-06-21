package remote

import (
	fluxerr "github.com/weaveworks/flux/errors"
)

func UnavailableError(err error) error {
	return &fluxerr.Error{
		Type: fluxerr.User,
		Help: `Cannot contact fluxd

To service this request, we need to ask the agent running in your
cluster (fluxd) for some information. But we can't connect to it at
present.

This may be because it's not running at all, or because it has
temporarily disconnected. You can check whether fluxd is connected with

    fluxctl status

If you are sure fluxd is running, you can simply wait a few seconds
and try the operation again.

If you are not sure if fluxd is running, please consult the
installation instructions in our documentation:

    https://github.com/weaveworks/flux/blob/master/site/installing.md

If you are still stuck, please log an issue:

    https://github.com/weaveworks/flux/issues

`,
		Err: err,
	}
}

func UpgradeNeededError(err error) error {
	return &fluxerr.Error{
		Type: fluxerr.User,
		Help: `Your fluxd needs to be upgraded

To service this request, we need to ask the agent running in your
cluster (fluxd) to perform an operation on our behalf, but the
version you have running is too old to understand the request.

Please install the latest version of fluxd and try again.

`,
		Err: err,
	}
}

func ClusterError(err error) error {
	return &fluxerr.Error{
		Type: fluxerr.User,
		Help: `Error from Flux daemon

The Flux daemon (fluxd) reported this error:

    ` + err.Error() + `

which indicates that it is running, but cannot complete the request.

Thus may be because the request wasn't valid; e.g., you asked for
something in a namespace that doesn't exist.

Otherwise, it is likely to be an ongoing problem until fluxd is
updated and/or redeployed. For help, please consult the installation
instructions:

    https://github.com/weaveworks/flux/blob/master/site/installing.md

If you are still stuck, please log an issue:

    https://github.com/weaveworks/flux/issues

`,
		Err: err,
	}
}
