export const fetchConfiguration = async () => {
    const res = await fetch('http://localhost:1323/api/config')
    return res.json()
}

export const saveConfiguration = async (conf: any) => {
    const res = await fetch('http://localhost:1323/api/config', {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(conf)
    })
    return res
}

export const fetchLogs = async (nItems: number) => {
    const res = await fetch(`http://localhost:1323/api/log?nItems=${nItems}`)
    let x = await res.text()
    return x;
}