# Segmentation Testing Scanner

This document outlines the workflow for performing segmentation testing using a combination of `nmap` and basic shell commands. Follow the steps below to set up and execute your scans efficiently.

## Idea

- Implement the scanner in Golang to enhance performance and concurrency.
- Utilize native Golang functions to scan the target.

## Workflow Steps

### 1. Set Target Network

Create a file named `cidr_network.txt` to specify the target CIDR network.

```bash
nano cidr_network.txt
```

### 2. Quick Network Scan

Perform a quick scan of the target network to identify live hosts. The results will be saved to `nst_sn_from_vlan10.txt`.

```bash
nmap -sn -iL cidr_network.txt | tee nst_sn_from_vlan10.txt
```

### 3. Extract Active IP Addresses

Extract only the IP addresses of the active hosts from the scan results and save them to `host_up_from_vlan10.txt`.

```bash
grep -o '[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}' nst_sn_from_vlan10.txt > host_up_from_vlan10.txt
```

### 4. Count Hosts That Are Up

Count the number of hosts that responded to the scan.

```bash
wc -l host_up_from_vlan10.txt
```

### 5. Scan TCP Ports

Conduct a TCP port scan on the identified live hosts. The results will be saved to `port_open_from_vlan10.txt`.

```bash
nmap -sT -iL host_up_from_vlan10.txt | tee port_open_from_vlan10.txt
```
