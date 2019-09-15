import React from "react";
import { Modal, Form, Button } from "react-bootstrap";

const PreviewModal = props => {
  console.log(props);
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
      <Modal.Body>
        <p>
          <b>Question:</b> {props.curr.question}
        </p>
        <p>
          <b>Answers:</b> {props.curr.answers}
        </p>
      </Modal.Body>
      <Modal.Footer>
        <Button onClick={props.onHide}>Close</Button>
      </Modal.Footer>
    </Modal>
  );
};

export default PreviewModal;
