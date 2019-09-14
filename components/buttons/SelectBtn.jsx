import React from "react";
import { Button } from "react-bootstrap";

const handleClick = () => {
  console.log("Selected!");
};

const SelectBtn = props => {
  return (
    <Button variant="success" onClick={handleClick}>
      Select
    </Button>
  );
};

export default SelectBtn;
