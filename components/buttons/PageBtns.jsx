import React from "react";
import { Button } from "react-bootstrap";

const PageBtns = () => {
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
      <Button
        onClick={() => {
          console.log("Finshed");
        }}
      >
        Finish
      </Button>
    </React.Fragment>
  );
};

export default PageBtns;
