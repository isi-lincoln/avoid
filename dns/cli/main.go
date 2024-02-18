package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	pkg "github.com/isi-lincoln/avoid/pkg"
	avoid "github.com/isi-lincoln/avoid/protocol"
)

var (
	dnsServer string
	dnsPort   int
	addr      string
)

func main() {

	root := &cobra.Command{
		Use:   "dnsctl",
		Short: "avoid's dns controller",
	}

	root.PersistentFlags().StringVarP(
		&dnsServer, "server", "s", "localhost", "avoid service address to use")
	root.PersistentFlags().IntVarP(
		&dnsPort, "port", "p", pkg.DefaultAvoidDNSPort, "avoid service port to use")

	addr = fmt.Sprintf("%s:%d", dnsServer, dnsPort)

	create := &cobra.Command{
		Use:   "create",
		Short: "Create dns entry",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			createEntryFunc(args[0])
		},
	}
	root.AddCommand(create)

	edit := &cobra.Command{
		Use:   "Edit",
		Short: "Edit dns entry",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			createEntryFunc(args[0])
		},
	}
	root.AddCommand(edit)

	del := &cobra.Command{
		Use:   "delete",
		Short: "Remove dns entry",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			delEntryFunc(args[0])
		},
	}
	root.AddCommand(del)

	showConfig := &cobra.Command{
		Use:   "show",
		Short: "Show system config",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			showConfigFunc(args[0])
		},
	}
	root.AddCommand(showConfig)

	listAvoid := &cobra.Command{
		Use:   "list",
		Short: "Show system config",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			listAvoidFunc()
		},
	}
	root.AddCommand(listAvoid)

	root.Execute()
}

// TODO: Split this into create 1 from cli
// and create 1/many from a file
func createEntryFunc(fi string) {

	fmt.Printf("test")

	eps, err := pkg.LoadDNSConfig(fi)
	if err != nil {
		log.Fatal(err)
	}

	ep := eps[0]

	addr := fmt.Sprintf("%s:%d", dnsServer, dnsPort)
	pkg.WithAvoidDNS(addr, func(c avoid.DNSClient) error {
		req := &avoid.EntryRequest{
			Entry: ep,
		}

		fmt.Printf("sent request: %v\n", req)
		resp, err := c.Update(context.TODO(), req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", resp)

		return nil
	})
}

func delEntryFunc(key string) {

	// TODO: sanitize key, dont allow ""
	if key == "" {
		return
	}

	pkg.WithAvoidDNS(addr, func(c avoid.DNSClient) error {
		req := &avoid.EntryRequest{
			Entry: &avoid.DNSEntry{
				Id: key,
			},
		}

		resp, err := c.Delete(context.TODO(), req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", resp)

		return nil
	})

}

func showConfigFunc(key string) {

	addr := fmt.Sprintf("%s:%d", dnsServer, dnsPort)
	pkg.WithAvoidDNS(addr, func(c avoid.DNSClient) error {

		resp, err := c.Show(context.TODO(), &avoid.ShowRequest{
			Key: key,
		})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Show:\n")

		// TODO: json pretty
		fmt.Printf("%+v\n", resp)

		return nil
	})
}

func listAvoidFunc() {

	addr := fmt.Sprintf("%s:%d", dnsServer, dnsPort)
	pkg.WithAvoidDNS(addr, func(c avoid.DNSClient) error {

		resp, err := c.List(context.TODO(), &avoid.ListRequest{})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", resp)

		return nil
	})
}
