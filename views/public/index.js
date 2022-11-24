function removeFromDb(accountId){
    fetch(`/accounts?accountId=${accountId}`, {method: "Delete"}).then(res =>{
        if (res.status === 200){
            window.location.pathname = "/accounts"
        }
    })
}

function updateDb(accountId) {
    let input = document.getElementById(accountId)
    let newitem = input.value
    fetch(`/accounts`, {
        method: "PUT",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'id': accountId,
            'username': newitem
        })
    }).then(res =>{
        if (res.status === 200){
            alert("Database updated")
            window.location.pathname = "/accounts"
        }
    })
}
