# Web Scraper CLI Application

A simple and customizable **Web Scraper** built with **Go**,
designed to extract content from websites and save it in various
formats (JSON, CSV, or plain text). This project demonstrates efficient
use of Go's `net/http`, `GoQuery`, and file handling capabilities.

## **Features**

- **Scrape HTML elements**: Extract text from tags like `h1`, `p`, `a`, etc.
- **Save in multiple formats**: Output scraped data as JSON, CSV, or plain text.
- **CLI arguments for customization**: Control the URL, tags, and
  output format from the command line.
- **Error handling**: Manage issues like invalid URLs, timeouts,
  or missing elements gracefully.
- **Scalable design**: Easily extendable to support advanced features
  like pagination, authentication, or proxy support.

## **Getting Started**

### Prerequisites

- Go 1.22 or higher installed on your system.
- Internet connection for fetching web pages.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/go-web-scraper.git cd go-web-scraper
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Build the application:

   ```bash
   go build -o webscraper
   ```

### Usage

Run the scraper with the following command:

```bash
./webscraper --url "https://example.com" --tags "h1,p" --output "output.json"
```

#### **Command-Line Options:**

| Option     | Description                                              | Default Value   |
| ---------- | -------------------------------------------------------- | --------------- |
| `--url`    | The URL of the webpage to scrape (required).             | None (required) |
| `--tags`   | Comma-separated list of HTML tags to extract.            | `h1,p`          |
| `--output` | File name to save the scraped data (supports JSON, CSV). | `output.json`   |

### Example

1. Scrape headings (`h1`) and links (`a`) from a webpage and save as CSV:

   ```bash
   ./webscraper --url "https://example.com" --tags "h1,a" --output "data.csv"
   ```

   Output (`data.csv`):

   ```css
   Welcome to Example Learn more Contact Us
   ```

2. Scrape paragraphs (`p`) from a blog and save as JSON:

   bash

   Copy code

   ```bash
   ./webscraper --url "https://blog.example.com" --tags "p" --output "blog.json"
   ```

   Output (`blog.json`):

   ```json
   [
     "This is the first paragraph of the blog.",
     "Another interesting point here.",
     "Conclusion goes here."
   ]
   ```

## **How It Works**

1. **Input Parsing:**

   - Accepts URL, tags, and output options via CLI arguments.
   - Validates input to ensure required parameters are provided.

2. **Fetching HTML:**

   - Uses `net/http` to send an HTTP GET request to the specified URL.
   - Handles response status codes and gracefully exits on errors.

3. **HTML Parsing:**

   - Leverages `GoQuery` to parse the HTML and extract elements
     based on user-specified tags.
   - Supports multiple tags and nested selections.

4. **Output Formatting:**

   - Saves data in JSON or CSV format using Go's `encoding/json`
     and `encoding/csv` libraries.
   - Ensures consistent formatting and file creation.

## **Project Structure**

bash

Copy code

```bash
go-web-scraper/
├── main.go            # Main application logic
├── fetcher.go         # Handles HTTP requests
├── parser.go          # Processes and parses HTML content
├── output.go          # Saves data in JSON or CSV format
├── go.mod             # Dependency file
└── README.md          # Project documentation
```

## **Future Enhancements**

- **Pagination Handling**: Automatically follow "Next" links to scrape multiple pages.
- **Authentication**: Support login and session cookies for protected content.
- **Proxy Support**: Use proxies to bypass restrictions or avoid IP bans.
- **Concurrency**: Parallelize scraping for faster performance.

## **Contributing**

Contributions are welcome! Feel free to:

- Report bugs.
- Suggest features.
- Submit pull requests.

## **License**

This project is licensed under the MIT License. See the `LICENSE` file for details.

## **Acknowledgments**

- [GoQuery](https://github.com/PuerkitoBio/goquery): HTML parsing library.
- [Colly](https://github.com/gocolly/colly): Inspiration for efficient web scraping.
- The Go community for excellent documentation and tutorials.
