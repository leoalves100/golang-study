// Ref: https://cli.urfave.org/v2

package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar retorna aplicação
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Artisan CLI"
	app.Usage = "Busca IPs e Nomes de Servidores na internete"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca ips",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscarIPs,
		},
		{
			Name:  "server",
			Usage: "Busca nome dos hosts",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscaDNS,
		},
	}

	return app
}

func buscarIPs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupHost(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscaDNS(c *cli.Context) {
	host := c.String("host")

	res, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ns := range res {
		fmt.Println(ns.Host)
	}

}
