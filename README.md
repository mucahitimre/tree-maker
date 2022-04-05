# TreeMaker
[//]: ===

## Of usage
`Converts a list with flat virtual paths to a tree data structure.`

### - Flat list
```
|—— "~/Foo/Videos/Clasic/museum.mp4"
|—— "~/Foo/Documents/doc/virtual.doc"
|—— "~/Foo/Documents/doc/real.doc"
|—— "~/Foo/Documents/doc/intuitive.doc"
|—— "~/Foo/Documents/doc/example.doc"
|—— "~/Foo/Documents/Pdf/analysis.pdf"
|—— "~/Bar/Exist/Pdf/analysis.pdf"
|—— "~/Mico/Sql/Pdf/analysis.pdf"
|—— "~/Data/Variable/Pdf/analysis.pdf" 
```
### - Tree list
```
[
  {
    "Key": "Foo",
    "Type": "Folder",
    "Childs": [
      {
        "Key": "Videos",
        "Type": "Folder",
        "Childs": [
          {
            "Key": "Clasic",
            "Type": "Folder",
            "Childs": [
              {
                "Key": "museum.mp4",
                "Type": "File",
                "Childs": []
              }
            ]
          }
        ]
      },
      {
        "Key": "Documents",
        "Type": "Folder",
        "Childs": [
          {
            "Key": "doc",
            "Type": "Folder",
            "Childs": [
              {
                "Key": "virtual.doc",
                "Type": "File",
                "Childs": []
              },
              {
                "Key": "real.doc",
                "Type": "File",
                "Childs": []
              },
              {
                "Key": "intuitive.doc",
                "Type": "File",
                "Childs": []
              },
              {
                "Key": "example.doc",
                "Type": "File",
                "Childs": []
              }
            ]
          },
          {
            "Key": "Pdf",
            "Type": "Folder",
            "Childs": [
              {
                "Key": "analysis.pdf",
                "Type": "File",
                "Childs": []
              }
            ]
          }
        ]
      }
    ]
  },
  {
    "Key": "Bar",
    "Type": "Folder",
    "Childs": [
      {
        "Key": "Exist",
        "Type": "Folder",
        "Childs": [
          {
            "Key": "Pdf",
            "Type": "Folder",
            "Childs": [
              {
                "Key": "analysis.pdf",
                "Type": "File",
                "Childs": []
              }
            ]
          }
        ]
      }
    ]
  },
  {
    "Key": "Mico",
    "Type": "Folder",
    "Childs": [
      {
        "Key": "Sql",
        "Type": "Folder",
        "Childs": [
          {
            "Key": "Pdf",
            "Type": "Folder",
            "Childs": [
              {
                "Key": "analysis.pdf",
                "Type": "File",
                "Childs": []
              }
            ]
          }
        ]
      }
    ]
  },
  {
    "Key": "Data",
    "Type": "Folder",
    "Childs": [
      {
        "Key": "Variable",
        "Type": "Folder",
        "Childs": [
          {
            "Key": "Pdf",
            "Type": "Folder",
            "Childs": [
              {
                "Key": "analysis.pdf",
                "Type": "File",
                "Childs": []
              }
            ]
          }
        ]
      }
    ]
  }
]
```
