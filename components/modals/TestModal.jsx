import React, { useState } from "react";
import { Modal, Form, Button } from "react-bootstrap";

const TestModal = props => {
  const [name, setName] = useState("");
  return (
    <Modal
      {...props}
      size="lg"
      aria-labelledby="contained-modal-title-vcenter"
      centered
    >
      <Modal.Header closeButton>
        <Modal.Title id="contained-modal-title-vcenter">
          {props.header}
        </Modal.Title>
      </Modal.Header>
      <Form>
        <Modal.Body>
          <Form.Group controlId="name">
            <Form.Label>Name</Form.Label>
            <Form.Control
              placeholder="Enter the test name"
              value={name}
              onChange={evt => setName(evt.target.value)}
            />
          </Form.Group>
        </Modal.Body>
      </Form>
      <Modal.Footer>
        <Button onClick={props.onHide}>Close</Button>
      </Modal.Footer>
    </Modal>
  );
};

export default TestModal;
