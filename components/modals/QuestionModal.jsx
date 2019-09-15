import React, { useState } from "react";
import { Modal, Form, Button } from "react-bootstrap";

const QuestionModal = props => {
  const [question, setQuestion] = useState("");
  const [answers, setAnswers] = useState("");

  const handleAdd = () => {
    console.log("QUESTION:    ", question);
    console.log("ANSWERS:    ", answers);
    props.onHide();
  };

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
          <Form.Group controlId="question">
            <Form.Label>Question</Form.Label>
            <Form.Control
              placeholder="Enter the question"
              value={question}
              onChange={evt => setQuestion(evt.target.value)}
            />
          </Form.Group>
          <Form.Group controlId="answers">
            <Form.Label>Answers</Form.Label>
            <Form.Control
              placeholder="Enter the answers"
              value={answers}
              onChange={evt => setAnswers(evt.target.value)}
            />
            <Form.Text>Enter the answers in a formatted way</Form.Text>
          </Form.Group>
        </Modal.Body>
      </Form>
      <Modal.Footer>
        <Button onClick={() => handleAdd(question, answers)}>Save</Button>
        <Button onClick={props.onHide}>Close</Button>
      </Modal.Footer>
    </Modal>
  );
};

export default QuestionModal;
