# Go Examples

Snippets of working Go examples.

## network-interfaces.go

    go run network-interfaces.go

Output:

	{
	  "en0": {
		"name": "en0",
		"macaddress": "00:00:00:a0:49:88",
		"ipaddress": [
		  "fe80::6555:35ff:fe51:1234/64",
		  "10.1.0.120/24"
		]
	  },
	  "gif0": {
		"name": "gif0",
		"macaddress": "",
		"ipaddress": null
	  },
	  "lo0": {
		"name": "lo0",
		"macaddress": "",
		"ipaddress": [
		  "fe80::1/64",
		  "127.0.0.1/8",
		  "::1/128"
		]
	  },
	  "p2p0": {
		"name": "p2p0",
		"macaddress": "00:00:00:a0:49:78",
		"ipaddress": null
	  },
	  "stf0": {
		"name": "stf0",
		"macaddress": "",
		"ipaddress": null
	  },
	  "vmnet1": {
		"name": "vmnet1",
		"macaddress": "00:00:00:a0:00:01",
		"ipaddress": [
		  "172.16.10.1/24"
		]
	  },
	  "vmnet8": {
		"name": "vmnet8",
		"macaddress": "00:00:00:a0:00:08",
		"ipaddress": [
		  "192.168.10.1/24"
		]
	  }
	}
    

## amazon-auth-header.go

    go run amazon-auth-header.go

Output:

    Authorization: AWS AKIAIOSFODNN7EXAMPLE:bWq2s1WEIj+Ydj0vQ697zp+IXMU=	

## readlines.go

Starting with go 1.1 the best way to read a file, line-by-line, is using a Scanner.

    printf "line1\nline2\nline3" > /tmp/test.txt
	go run readlines.go

Output:

    line1
	line2
	line3
