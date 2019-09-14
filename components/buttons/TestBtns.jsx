import React from "react";
import { Button } from "react-bootstrap";

const TestBtns = () => {
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

export default TestBtns;
