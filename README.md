# Quiver
Quiver is a library that facilitates queries on MongoDB Documents.

## Document
{
    "carsBrandArray": [
        {
            "brand": "BMW",
            "name": "i8",
            "value": 2000,
        },
        {
            "brand": "Chevy",
            "name": "Camaro",
            "value": 400,
        },
        {
            "brand": "Chevy",
            "name": "Malibu",
            "value": 300,
        },
        {
            "brand": "BMW",
            "name": "i3",
            "value": 400,
        },
    ]    
}

## Example
quiver.Query(
    quiver.AND(
        "carsBrandArray => brand == BMW && value >= 500",
        "carsBrandArray => brand == Chevy && value <= 350",
    ),
)

## MongoDB Query
{
    "$and": [
        {
            "carsBrandArray": {
                "$elemMatch": {
                    "name": "BMW",
                    "value": {
                        "$gte": 500
                    }
                }
            }
        },
        {
            "carsBrandArray": {
                "$elemMatch": {
                    "name": "Chevy",
                    "value": {
                        "$lte": 350
                    }
                }
            }
        }
    ]
}

## Results
BMW i8
Chevy Malibu