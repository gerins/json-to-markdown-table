# JSON to Markdown Table Converter

This repository provides a Go program to convert JSON data into Markdown tables. The tool dynamically generates markdown tables for any JSON structure, making it easy to document complex data schemas.

## Features

- Automatically parses JSON data and outputs nested markdown tables.
- Supports different data types including objects, arrays, strings, numbers, and booleans.
- Infers field requirements (`Mandatory` or `Optional`) based on JSON content.

## Example

### Input JSON
```json
{
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "is_active": true,
    "roles": ["admin", "user"]
  }
}
```

### Output Markdown
```markdown
### Main Structure
| Parameter | Required | Data Type    | Description      | Example              |
|-----------|----------|--------------|------------------|----------------------|
| user      | M        | Object       | Field for 'user' | Refer to sub-structure |

### user Structure
| Parameter  | Required | Data Type       | Description          | Example          |
|------------|----------|-----------------|----------------------|------------------|
| id         | M        | Number          | Field for 'id'       | 1                |
| name       | M        | String          | Field for 'name'     | John Doe         |
| email      | M        | String          | Field for 'email'    | john@example.com |
| is_active  | M        | Boolean         | Field for 'is_active'| true             |
| roles      | M        | Array String    | Field for 'roles'    | Refer to list    |

```

## Usage

1. Clone the repository:
   ```sh
   git clone https://github.com/gerins/json-to-markdown-table.git
   cd json-to-markdown-table
   ```

2. Run the program:
   ```sh
   go run main.go
   ```

3. Replace the `inputJSON` variable in `main.go` with your JSON data to generate a markdown table for your schema.

## Implementation

The core functionality is in the `GenerateMarkdown` function, which parses the JSON input and generates markdown tables. It uses recursion to handle nested objects and arrays dynamically.

### Key Functions

- `GenerateMarkdown`: Orchestrates the JSON parsing and markdown generation.
- `generateObjectTable`: Recursively generates markdown tables for JSON objects.
- `generateArrayTable`: Processes arrays of objects and generates the appropriate markdown.
- `getTypeAndExample`: Determines data types and examples for fields.
- `inferRequired`: Infers whether a field is `Mandatory` or `Optional`.

## 🙌 Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you'd like to change.


## 🔒 License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT). See the `LICENSE` file for details.


## ✍️ Author

**Garin Prakoso** 
[GitHub](https://github.com/gerins) | [LinkedIn](https://www.linkedin.com/in/garin-prakoso-60244b1a2/)
Feel free to contact me if you need help or have any feedback.

