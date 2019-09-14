import React from "react";
import { Modal, Form, Button } from "react-bootstrap";

const handleUpdate = () => {
  console.log("Updated");
};

const QuestionModal = props => {
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
        <Modal.Body>Moodal</Modal.Body>
      </Form>
      <Modal.Footer>
        <Button onClick={handleUpdate} variant="success">
          Update
        </Button>
        <Button onClick={props.onHide}>Close</Button>
      </Modal.Footer>
    </Modal>
  );
};

export default QuestionModal;
