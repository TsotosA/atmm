import React from "react";
import {Container, Nav, Navbar} from "react-bootstrap";
import {Link} from "react-router-dom";

function Head() {
    return (
        <>
            <Navbar fixed="bottom" bg="dark" variant="dark" expand="md">
                <Container>
                    <Navbar.Brand as={Link} to="/">@ media manager</Navbar.Brand>
                    <Navbar.Toggle aria-controls="basic-navbar-nav"/>
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Nav className="me-auto">
                            {/*<Nav.Link as={Link} to="/config">Configuration</Nav.Link>*/}
                            {/*<Nav.Link as={Link} to="/log">Logs</Nav.Link>*/}
                        </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
        </>
    );
}

export default Head;
