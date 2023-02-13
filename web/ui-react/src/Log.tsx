import {useQuery} from "@tanstack/react-query";
import {fetchLogs} from "./service";
import {InputHTMLAttributes, useEffect, useState} from "react";
// @ts-ignore
import _ from 'lodash';

export function Log() {
    const [nItems, setNItems] = useState(100);
    const [grepFor, setGrepFor] = useState('');
    const {data, isLoading} = useQuery(['logs', nItems, grepFor], () => {
        return fetchLogs(nItems, grepFor)
    })

    let handleLinesInput = (value: any) => {
        if (+value > 10000) {
            return
        }
        setNItems(+value)
    }

    let handleGrepInput = (value: any) => {
        console.log('grep input')
        setGrepFor(value)
    }

    // A debounced input react component
    function DebouncedInput({
                                value: initialValue,
                                onChange,
                                debounce = 500,
                                ...props
                            }: {
        value: string | number
        onChange: (value: string | number) => void
        debounce?: number
    } & Omit<InputHTMLAttributes<HTMLInputElement>, 'onChange'>) {
        const [value, setValue] = useState(initialValue)

        useEffect(() => {
            setValue(initialValue)
        }, [initialValue])

        useEffect(() => {
            const timeout = setTimeout(() => {
                onChange(value)
            }, debounce)

            return () => clearTimeout(timeout)
        }, [value])

        return (
            <input {...props} value={value} onChange={e => setValue(e.target.value)}/>
        )
    }

    return (
        <>
            <div className={"shadow1 p-3 mb-5 rounded"}>
                <h1 className="display-6">Logs</h1>
                <form>
                    <div className="form-group col-lg-2 col-sm-2">
                        <label htmlFor="numberOfLinesInput">Number of lines</label>s
                        {/*<input onChange={handleLinesInput} type="text" className="form-control form-control-sm"*/}
                        {/*       id="numberOfLinesInput"*/}
                        {/*       placeholder="ex. 200" value={nItems} style={{backgroundColor: "dimgrey"}}></input>*/}
                        <DebouncedInput value={nItems} onChange={value => handleLinesInput(+value)}
                                        debounce={1000} id="numberOfLinesInput" className="form-control form-control-sm"
                                        style={{backgroundColor: "dimgrey"}} placeholder="ex. 200"/>
                    </div>
                    <small id="emailHelp" className="form-text text-muted">enter a number for lines to tail from
                        logs. 0 fetches the whole file.</small>
                    <div className="form-group col-lg-4 col-sm-4">
                        <label htmlFor="grepInput">grep</label>
                        {/*<input onChange={handleGrepInput} type="text" className="form-control form-control-sm"*/}
                        {/*       id="grepInput"*/}
                        {/*       placeholder="ex. error" value={grepFor} style={{backgroundColor: "dimgrey"}}></input>*/}
                        <DebouncedInput value={grepFor} onChange={value => handleGrepInput(value)}
                                        debounce={1000} id="grepInput" className="form-control form-control-sm"
                                        style={{backgroundColor: "dimgrey"}} placeholder="ex. error"/>
                    </div>
                    <small id="emailHelp" className="form-text text-muted">enter a string to search for in the selected
                        lines</small>

                </form>
                <pre>{isLoading ? 'loading' : data ? data : 'no data'}</pre>
            </div>
        </>
    );
}