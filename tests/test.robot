*** Settings ***
Library         Collections
Library         RequestsLibrary
Library         REST
Library         BuiltIn

*** Variables ***
${base_url} =       http://localhost:8080

*** Test Cases ***
Get item
    ${id}               Evaluate            random.randint(100, 300)        modules=random
    ${sub_key_id}       Evaluate            random.randint(1, 3)            modules=random
    ${url}              Set Variable        ${base_url}/item/${id}/S-${sub_key_id}
    Log To Console      URL: ${url}
    GET                 ${url}
    Integer             response status     200
    Log                 response body

Get item NOT existing
    ${id}               Set Variable        100000
    ${sub_key_id}       Set Variable        100
    ${url}              Set Variable        ${base_url}/item/${id}/S-${sub_key_id}
    Log To Console      URL: ${url}
    GET                 ${url}
    Integer             response status     404

Get Items for ID
    ${id}               Evaluate            random.randint(100, 300)        modules=random
    ${url}              Set Variable        ${base_url}/items/${id}
    Log To Console      URL: ${url}
    GET                 ${url}
    Integer             response status     200
    Log                 response body
    Integer             $.Count             3

Get Items
    ${url}              Set Variable        ${base_url}/items
    Log To Console      URL: ${url}
    GET                 ${url}
    Integer             response status     200

Get Items with max
    ${max}              Evaluate            random.randint(1, 5)                modules=random
    ${url}              Set Variable        ${base_url}/items?max=${max}
    Log To Console      URL: ${url}
    GET                 ${url}
    Integer             response status     200
    Log                 response body
    Integer             $.Count             ${max}  

Get Items with max (negative max) - error case
    ${max}              Set Variable        -10    
    ${url}              Set Variable        ${base_url}/items?max=${max}
    Log To Console      URL: ${url}
    GET                 ${url}
    Integer             response status     400

Get Items with max (alpha max) - error case
    ${max}              Set Variable        abcd    
    ${url}              Set Variable        ${base_url}/items?max=${max}
    Log To Console      URL: ${url}
    GET                 ${url}
    Integer             response status     400

Get Items with max (exceeding MAX) - error case
    ${max}              Set Variable        10000   
    ${url}              Set Variable        ${base_url}/items?max=${max}
    Log To Console      URL: ${url}
    GET                 ${url}
    Integer             response status     400

Not handled uri - error case
    ${url}              Set Variable        ${base_url}/wrong
    Log To Console      URL: ${url}
    GET                 ${url}
    Integer             response status     404