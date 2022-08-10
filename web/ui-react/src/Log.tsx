import {useQuery} from "@tanstack/react-query";
import {fetchLogs} from "./service";
import {useState} from "react";

export function Log() {
    const [nItems, setNItems] = useState(20);
    const {data} = useQuery(['logs', nItems], () => {
        return fetchLogs(nItems)
    })

    let handleInput = ($event: any) => {
        if (+$event.target.value > 10000) {
            return
        }
        setNItems($event.target.value)
    }

    return (
        <>
            <div className={"shadow1 p-3 mb-5 rounded"}>
                <h1 className="display-6">Logs</h1>
                <form>
                    <div className="form-group col-lg-2 col-sm-2">
                        <label htmlFor="numberOfLinesInput">Number of lines</label>
                        <input onChange={handleInput} type="text" className="form-control form-control-sm"
                               id="numberOfLinesInput"
                               aria-describedby="emailHelp"
                               placeholder="ex. 200" value={nItems} style={{backgroundColor: "dimgrey"}}></input>
                        <small id="emailHelp" className="form-text text-muted">enter a number for lines to tail from
                            logs</small>
                    </div>
                </form>
                <pre>{data ? data : "loading"}</pre>
            </div>
        </>
    );
}