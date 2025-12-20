#!/usr/bin/env python3
"""Agent Service - Main entry point."""

import os
from http.server import HTTPServer, BaseHTTPRequestHandler


class AgentHandler(BaseHTTPRequestHandler):
    """HTTP request handler for Agent service."""

    def do_GET(self):
        """Handle GET requests."""
        self.send_response(200)
        self.send_header("Content-type", "text/plain")
        self.end_headers()
        self.wfile.write(b"Agent Service is running\n")

    def log_message(self, format, *args):
        """Override to customize logging."""
        return


def main():
    """Main function to start the server."""
    port = int(os.getenv("PORT", "8000"))
    server_address = ("", port)
    httpd = HTTPServer(server_address, AgentHandler)
    print(f"Agent Service starting on port {port}")
    httpd.serve_forever()


if __name__ == "__main__":
    main()

