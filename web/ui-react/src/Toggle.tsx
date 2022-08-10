import Accordion from 'react-bootstrap/Accordion';
import {useAccordionButton} from 'react-bootstrap/AccordionButton';
import Card from 'react-bootstrap/Card';
import {Button} from "react-bootstrap";

// @ts-ignore
export function CustomToggle({children, eventKey}) {
    const decoratedOnClick = useAccordionButton(eventKey, () =>{}
        // console.log('totally custom!'),
    );

    return (
        <>
            {/*<Button onClick={decoratedOnClick} variant="secondary"> {children}</Button>*/}

            <svg onClick={decoratedOnClick} xmlns="http://www.w3.org/2000/svg" width="26" height="26" fill="currentColor"
                 className="bi bi-caret-down" viewBox="0 0 16 16">
                <path
                    d="M3.204 5h9.592L8 10.481 3.204 5zm-.753.659 4.796 5.48a1 1 0 0 0 1.506 0l4.796-5.48c.566-.647.106-1.659-.753-1.659H3.204a1 1 0 0 0-.753 1.659z"/>
            </svg>

            <svg onClick={decoratedOnClick} xmlns="http://www.w3.org/2000/svg" width="26" height="26" fill="currentColor"
                 className="bi bi-caret-up" viewBox="0 0 16 16">
                <path
                    d="M3.204 11h9.592L8 5.519 3.204 11zm-.753-.659 4.796-5.48a1 1 0 0 1 1.506 0l4.796 5.48c.566.647.106 1.659-.753 1.659H3.204a1 1 0 0 1-.753-1.659z"/>
            </svg>

        </>
    );
}

// @ts-ignore
export function Toggle({children}) {
    return (
        <Accordion defaultActiveKey="1" className={"w-100"}>
            <Card >
                <Card.Header>
                    <CustomToggle eventKey="0">config.yaml</CustomToggle>
                </Card.Header>
                <Accordion.Collapse eventKey="0" >
                    <Card.Body >{children}</Card.Body>
                </Accordion.Collapse>
            </Card>
        </Accordion>
    );
}

