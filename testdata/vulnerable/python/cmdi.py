import os
import sys

def ping_host(host):
    """
    Pings a host to check if it's alive.
    """
    print(f"Pinging {host}...")
    
    # VULNERABLE: Direct OS command injection
    os.system("ping -c 4 " + host)

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python cmdi.py <host>")
        sys.exit(1)
        
    ping_host(sys.argv[1])
