#!/usr/bin/env python3
import requests

BASE = "https://api.terminjetzt.com"

def latest():
    return requests.get(f"{BASE}/appointments/latest").json()

if __name__ == "__main__":
    print(latest())
