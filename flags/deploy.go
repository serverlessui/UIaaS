package flags

import "github.com/urfave/cli"

const (
	hostedZoneArg    = "hostedzone, ho"
	domainNameArg    = "domainname, d"
	cacheTTL         = "cachettl, c"
	hostedZoneExists = "hostedzoneexists, e"
	tag              = "tag, t"
	environment      = "environment, env"
	appdir           = "applicationdirectory, dir"
)

//Deploy method to return flags for deployment
func Deploy() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  hostedZoneArg,
			Usage: "Route 53 `hostedzone`",
		},
		cli.StringFlag{
			Name:  domainNameArg,
			Usage: "`dns name` for serverless ui application",
		},
		cli.StringFlag{
			Name:  cacheTTL,
			Usage: "`cache ttl` for Cloudfront cache",
		},
		cli.StringFlag{
			Name:  hostedZoneExists,
			Usage: "Route 53 `hostedzone exists`",
		},
		cli.StringFlag{
			Name:  tag,
			Usage: "`tag` used to tag resources for tracking and billing ",
		},
		cli.StringFlag{
			Name:  environment,
			Usage: "`environment` used to differentiate deployments ",
		},
		cli.StringFlag{
			Name:  appdir,
			Usage: "`applicationdirectory` Directory containing ui source code",
		},
	}
}
