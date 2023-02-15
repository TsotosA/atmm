export const fetchConfiguration = async () => {
    const res = await fetch(`${process.env.REACT_APP_API_ROOT}/api/config`)
    return res.json()
}

export const saveConfiguration = async (conf: any) => {
    const res = await fetch(`${process.env.REACT_APP_API_ROOT}/api/config`, {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(conf)
    })
    return res
}

export const fetchLogs = async (nItems: number, grepFor?: string) => {
    const res = await fetch(`${process.env.REACT_APP_API_ROOT}/api/log?nItems=${nItems}&grepFor=${grepFor}`)
    return await res.text();
}