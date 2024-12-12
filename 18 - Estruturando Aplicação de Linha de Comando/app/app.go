package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Domínios"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de Domínios",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
					Usage: "Nome do host para buscar os IPs",
				},
			},
			Action: buscarIPs,
		},
		{
			Name:  "servidores",
			Usage: "Busca Nomes de Domínios (NS)",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
					Usage: "Nome do host para buscar os domínios",
				},
			},
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIPs(context *cli.Context) error {
	host := context.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatalf("Erro ao buscar IPs: %v", err)
	}

	fmt.Printf("IPs para o host %s:\n", host)
	for _, ip := range ips {
		fmt.Println(ip)
	}

	return nil
}

func buscarServidores(context *cli.Context) error {
	host := context.String("host")

	servidores, err := net.LookupNS(host)
	if err != nil {
		log.Fatalf("Erro ao buscar servidores de nomes: %v", err)
	}

	fmt.Printf("Servidores de nomes para o host %s:\n", host)
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

	return nil
}