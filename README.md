# oracle-code-berlin
Demo fn project functions from my Oracle Code Berlin presentation.

# Project folders
Project/Folder | Description
-------------- | ------------
test_data | Folder containing two test data files. Files contain JSON booking records.
searchProduct | Simple Go function extracting all booking records for a given product name from input. File test_data/MOCK_DATA_SEARCH.json contains valid records.
extractEmail | Simpel Go function extracting all email addresses from an JSON Array containing booking records. Function searchProduct provides valid data.
identifyMarketingTargetsFlow | Fn flow function coordinating searchProduct and extractEmail functinos. File test_data/MOCK_DATA_SEARCH.json contains valid records.

## Assumptions
In order to execute the following commands the following assumptions are made:
1. All three functions were deployed under an application called booking
2. fn server runs locally on port 8080
3. curl commands are executed from within the repository root directory

## searchProduct function
The search product function can be called as follows:

```
curl -X POST -d @./test_data/MOCK_DATA_SEARCH.json http://localhost:8080/r/booking/searchProduct
```

## extractEmail function
The extract email function expects searchProduct function's output as input. Within the following command it is assumed the output was saved in a file called IDENTIFIED_BOOKINGS.json.

The extract email function can be called as follows:

```
curl -X POST -d @./IDENTIFIED_BOOKINGS.json http://localhost:8080/r/booking/extractEmail
```

Alternatively searchProduct and extractEmail functions can be chained on the command line:

```
curl -X POST -d @./test_data/MOCK_DATA_SEARCH.json http://localhost:8080/r/booking/searchProduct | curl -X POST -d @- http://localhost:8080/r/booking/extractEmail
```

## identifyMarketingTargetsFlow function
The identifyMarketingTargetsFlow function can be called as follows:

```
curl -X POST -d @./test_data/MOCK_DATA_SEARCH.json http://localhost:8080/r/booking/identifyMarketingTargetsFlow
```
