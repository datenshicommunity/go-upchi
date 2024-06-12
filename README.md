# Go-Upchi: Temporary file uploader

Go-Upchi is a lightweight and secure file-sharing solution built in Golang, designed specifically for internal network. It will be secure to share without any problem into your teams.

## Demo
<iframe width="560" height="315" src="https://www.youtube.com/embed/_O6Bv45z08g?si=SU3C39zHZsEEZYdX" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>

## Key Features

- **Universal Compatibility:** Compatible with any S3 server, including R2 Cloudflare, Amazon S3 and MinIO.
- **Temporary Sharing:** Share files with predefined expiration periods, reducing clutter and promoting efficient file management.
- **Intuitive User Interface:** Simple and intuitive interface for easy file upload, and sharing.

## Getting Started

### Prerequisites

- Go (1.20-latest)
- S3 Server (R2, AWS, MiniO)

### Installation

Clone the repository:

```bash
git clone https://github.com/datenshicommunity/go-upchi.git
```

Build the project:

```bash
cd go-upchi
go build
```

Config the environment:

```
cp .env.example .env
```

Run the executable:

```bash
./go-upchi
```

Visit `http://localhost:3000` in your web browser to access the application.

## Usage

When opening the application on the browsers, you can drag n drop the file on there and then upload the file wait until link url show.

## Contributing

We welcome contributions from the community!

## License

This project is licensed under the [MIT License](LICENSE).