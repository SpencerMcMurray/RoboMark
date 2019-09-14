import React from "react";
import { Button } from "react-bootstrap";

const QuestionBtns = () => {
  return (
    <React.Fragment>
      <Button
        variant="success"
        onClick={() => {
          console.log("Created new");
        }}
      >
        Create New
      </Button>
    </React.Fragment>
  );
};

export default QuestionBtns;
