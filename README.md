# Gmap

A (disturbingly fast) port scanner written in Go. 

## Usage

```bash
./Gmap -t scanme.nmap.org -T 250 -p 1024
```

## Notes

Higher thread amounts will significantly increase speed, but results will decrease in accuracy. 
