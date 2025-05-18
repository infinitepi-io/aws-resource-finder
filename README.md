# AWS Resource Finder

A web-based tool built in Go that helps you discover and visualize AWS resources across your AWS account using AWS Config.

![Infinite Pi Logo](template/infinite-pi.png)

## Features

- 🔍 Search AWS resources by resource type
- 🌐 Web-based interface for easy access
- 📊 Table view of resource details
- ⚡ Real-time resource querying using AWS Config
- 🔄 Support for multiple AWS resource types (Lambda, S3, EC2, etc.)

## Prerequisites

- Go 1.16 or higher
- AWS Account with AWS Config enabled
- AWS CLI configured with appropriate credentials
- AWS SDK for Go

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/aws-resource-finder.git
cd aws-resource-finder
```

2. Install dependencies:
```bash
go mod download
```

## Configuration

Ensure you have AWS credentials configured either through:
- AWS CLI (`aws configure`)
- Environment variables
- IAM role

The application needs permissions to:
- Query AWS Config
- List resources

## Usage

1. Start the server:
```bash
go run main.go
```

2. Open your web browser and navigate to:
```
http://localhost:8080/aws-resource-finder
```

3. Select a resource type from the dropdown and click "Find Resources"

## API Query Format

The application uses AWS Config query syntax:
```sql
SELECT * WHERE resourceType = '<AWS::Service::Resource>'
```

Supported resource types include:
- AWS::Lambda::Function
- AWS::S3::Bucket
- AWS::EC2::Instance

## Project Structure

```
├── main.go           # Main application entry point
├── aws/
│   └── aws.go       # AWS Config integration
├── template/        # HTML templates and static assets
│   ├── result.html
│   └── images/
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- AWS Config Documentation
- Go Web Programming Community
- Contributors and maintainers

## Support

If you encounter any issues or have questions, please:
1. Check the existing issues
2. Create a new issue with details about your problem
3. Include relevant logs and configuration