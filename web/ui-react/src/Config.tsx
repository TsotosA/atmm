import {useQuery} from "@tanstack/react-query";
import {fetchConfiguration, saveConfiguration} from "./service";
import {Button, Table} from "react-bootstrap";
import {Toggle} from "./Toggle";
import {useState} from "react";

export function Config() {

    const {data, status} = useQuery(['config'], fetchConfiguration)
    const [newConf, setNewConf] = useState(data)

    // useEffect(() => {
    //     setNewConf(data)
    // }, [data])

    const valueChange = (e: { key: string, val: any }) => {
        let tmp = {...newConf};
        tmp[e.key] = e.val
        setNewConf({...tmp})
    }

    function save() {
        console.log('save')
        console.log(newConf)
        saveConfiguration(newConf).then(r => console.log('returned from save'))
    }

    return (
        status !== 'success' ? <>Loading</> :
            <>
                <div className={"shadow1 p-3 mb-5 rounded"}>
                    <h1 className="display-6">Configuration</h1>
                    <Toggle>
                        <Table className={"w-auto"} size="sm" responsive striped bordered hover variant="dark">
                            <thead>
                            <tr>
                                <th>Key</th>
                                <th>Current value</th>
                            </tr>
                            </thead>
                            <tbody>
                            {
                                data && Object.keys(data).map(key => {
                                    return (
                                        <tr key={key}>
                                            <td className="fw-semibold">{key}</td>
                                            <td className="fw-lighter">
                                                {data[key].toString() || 'no value'}
                                                {/*<span style={{color: "red"}}>*/}
                                                {/*   &nbsp;--&gt;&nbsp;{data[key].toString() || 'no value'}*/}
                                                {/*</span>*/}
                                            </td>
                                            <td className="fw-lighter">
                                                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20"
                                                     fill="currentColor" className="bi bi-pencil-square"
                                                     viewBox="0 0 16 16">
                                                    <path
                                                        d="M15.502 1.94a.5.5 0 0 1 0 .706L14.459 3.69l-2-2L13.502.646a.5.5 0 0 1 .707 0l1.293 1.293zm-1.75 2.456-2-2L4.939 9.21a.5.5 0 0 0-.121.196l-.805 2.414a.25.25 0 0 0 .316.316l2.414-.805a.5.5 0 0 0 .196-.12l6.813-6.814z"/>
                                                    <path fill-rule="evenodd"
                                                          d="M1 13.5A1.5 1.5 0 0 0 2.5 15h11a1.5 1.5 0 0 0 1.5-1.5v-6a.5.5 0 0 0-1 0v6a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-11a.5.5 0 0 1 .5-.5H9a.5.5 0 0 0 0-1H2.5A1.5 1.5 0 0 0 1 2.5v11z"/>
                                                </svg>
                                                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20"
                                                     fill="green" className="bi bi-check2-square"
                                                     viewBox="0 0 16 16">
                                                    <path
                                                        d="M3 14.5A1.5 1.5 0 0 1 1.5 13V3A1.5 1.5 0 0 1 3 1.5h8a.5.5 0 0 1 0 1H3a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5V8a.5.5 0 0 1 1 0v5a1.5 1.5 0 0 1-1.5 1.5H3z"/>
                                                    <path
                                                        d="m8.354 10.354 7-7a.5.5 0 0 0-.708-.708L8 9.293 5.354 6.646a.5.5 0 1 0-.708.708l3 3a.5.5 0 0 0 .708 0z"/>
                                                </svg>
                                                <input type="text" className="form-control form-control-sm"
                                                       id="numberOfLinesInput"
                                                       aria-describedby="emailHelp"
                                                       placeholder="ex. 200" value={'nItems'} style={{backgroundColor: "dimgrey", width:'150px', display:"inline"}}>

                                                </input>
                                                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20"
                                                     fill="red" className="bi bi-x-square" viewBox="0 0 16 16">
                                                    <path
                                                        d="M14 1a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h12zM2 0a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2H2z"/>
                                                    <path
                                                        d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708z"/>
                                                </svg>
                                            </td>
                                        </tr>
                                    )
                                })
                            }
                            </tbody>
                        </Table>
                        {/*<Button variant="outline-success" onClick={save}>Save conf changes</Button>*/}
                    </Toggle>
                </div>
            </>
    )
}