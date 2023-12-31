# Kaffeine

**This project is in Beta. Updates will have to be done manually for now.**

![Kaffeine Logo](https://github.com/BolajiOlajide/kaffeine/blob/main/assets/screenshot.png?raw=true)

## Description

Kaffeine is a lightweight, efficient, and reliable library designed to keep your web services active and alert! In the cloud-driven era, several services can get spun down due to inactivity. Whether you're running a critical service or a side project, Kaffeine ensures that your URLs stay awake, online, and responsive, warding off any unintended inactivity timeouts set by cloud providers.

## Features

- **Multi-URL Support**: Ping multiple URLs simultaneously without any hassle.
- **Configurable Intervals**: Set custom intervals for each URL or use default settings.
- **Minimal Overhead**: Efficiently designed to use minimal resources.
- **Logs and Reports**: Get detailed logs and reports for each ping, helping you diagnose any potential issues.

## Installation

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/BolajiOlajide/kaffeine/main/install.sh)"
```

## Usage

Here's a quick start guide to get Kaffeine running:

```sh
kaffeine -url https://www.example.com -interval 60
```

## Contributing

We welcome contributions! Whether it's bug reports, feature requests, or pull requests – all are appreciated and reviewed. Please check the [CONTRIBUTING.md](./CONTRIBUTING.md) for more details.

## License

Kaffeine is released under the [MIT License](./LICENSE).

Stay awake with Kaffeine! ☕
