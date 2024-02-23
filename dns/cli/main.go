package main

import (
	"context"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	pkg "github.com/isi-lincoln/avoid/pkg"
	avoid "github.com/isi-lincoln/avoid/protocol"
)

var (
	dnsServer   string
	dnsPort     int
	addr        string
	arecords    []string
	aaaarecords []string
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

	update := &cobra.Command{
		Use:   "update",
		Short: "create or modify dns entry",
	}
	root.AddCommand(update)

	updateCli := &cobra.Command{
		Use:   "cli",
		Short: "update dns entry from cli",
		Long:  "cli <ue> <name> <ttl> (<txt>) --a 192.168.0.1  --aaaa dead:beef::0123 ...",
		Args:  cobra.MinimumNArgs(2), // id, record type, ttl. then use flags for records
		Run: func(cmd *cobra.Command, args []string) {

			ue := args[0]
			name := args[1]
			ttlStr := args[2]
			txt := ""
			if len(args) > 3 {
				txt = args[3]
			}
			// Id: for now ip address
			// Recordtype: for now either A/AAAA
			// Records: dns names associated
			// TTL: Time to Live for records
			// TXT: additional field that can be used for validation
			arecordsArr, err := cmd.Flags().GetStringArray("a")
			if err != nil {
				fmt.Printf("Error parsing flag 'a': %v\n", err)
				return
			}

			aaaarecordsArr, err := cmd.Flags().GetStringArray("aaaa")
			if err != nil {
				fmt.Printf("Error parsing flag 'aaaa': %v\n", err)
				return
			}

			ttl, err := strconv.Atoi(ttlStr)
			if err != nil {
				fmt.Printf("Unable to convert TTL to Int: %v\n", err)
				return
			}

			entry := &avoid.DNSEntry{
				Ue:          ue,
				Name:        name,
				Ttl:         int64(ttl),
				Txt:         txt,
				Arecords:    arecordsArr,
				Aaaarecords: aaaarecordsArr,
			}

			err = pkg.CheckDNSRecord(entry)
			if err != nil {
				fmt.Printf("Invalid entry: %v\n", err)
				return
			}

			updateCliFunc(entry)
		},
	}
	update.AddCommand(updateCli)

	updateCli.Flags().StringArrayVar(&arecords, "a", []string{}, "DNS Entry record value for A records")
	updateCli.Flags().StringArrayVar(&aaaarecords, "aaaa", []string{}, "DNS Entry record value for AAAA records")

	updateFile := &cobra.Command{
		Use:   "file",
		Short: "update dns entries from file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			updateFileFunc(args[0])
		},
	}
	update.AddCommand(updateFile)

	del := &cobra.Command{
		Use:   "delete",
		Short: "Remove dns entry",
		Long:  "delete <ue> <name>",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			delEntryFunc(args[0], args[1])
		},
	}
	root.AddCommand(del)

	showConfig := &cobra.Command{
		Use:   "show",
		Short: "Show system config",
		Long:  "show <ue> <name>",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			showConfigFunc(args[0], args[1])
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

func updateCliFunc(entry *avoid.DNSEntry) {

	fmt.Printf("test cli\n")

	addr := fmt.Sprintf("%s:%d", dnsServer, dnsPort)
	pkg.WithAvoidDNS(addr, func(c avoid.AVOIDDNSClient) error {
		req := &avoid.EntryRequest{
			Entries: []*avoid.DNSEntry{entry},
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

func updateFileFunc(fi string) {

	fmt.Printf("test file")

	eps, err := pkg.LoadDNSConfig(fi)
	if err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf("%s:%d", dnsServer, dnsPort)
	pkg.WithAvoidDNS(addr, func(c avoid.AVOIDDNSClient) error {
		req := &avoid.EntryRequest{
			Entries: eps,
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

func delEntryFunc(ue, name string) {

	// TODO: sanitize key, dont allow ""
	if ue == "" || name == "" {
		fmt.Printf("invalid key to delete: %s/%s", ue, name)
		return
	}

	entry := &avoid.DNSEntry{
		Ue:   ue,
		Name: name,
	}

	pkg.WithAvoidDNS(addr, func(c avoid.AVOIDDNSClient) error {
		req := &avoid.EntryRequest{
			Entries: []*avoid.DNSEntry{entry},
		}

		resp, err := c.Delete(context.TODO(), req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", resp)

		return nil
	})

}

func showConfigFunc(ue, name string) {

	addr := fmt.Sprintf("%s:%d", dnsServer, dnsPort)
	pkg.WithAvoidDNS(addr, func(c avoid.AVOIDDNSClient) error {

		resp, err := c.Show(context.TODO(), &avoid.ShowRequest{
			Ue:   ue,
			Name: name,
		})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Show:\n")

		fmt.Printf("%+v\n", resp)

		return nil
	})
}

func listAvoidFunc() {

	var keys []string
	addr := fmt.Sprintf("%s:%d", dnsServer, dnsPort)
	pkg.WithAvoidDNS(addr, func(c avoid.AVOIDDNSClient) error {

		resp, err := c.List(context.TODO(), &avoid.ListRequest{})

		if err != nil {
			log.Fatal(err)
		}

		keys = resp.Keys

		return nil
	})

	if len(keys) > 0 {
		fmt.Printf("Keys:\n")
		for _, k := range keys {
			fmt.Printf("\t%s\n", k)
		}
	} else {
		fmt.Printf("No dns keys found\n")
	}
}
