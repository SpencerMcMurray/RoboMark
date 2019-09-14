import React from "react";
import { Button } from "react-bootstrap";

const handleClick = () => {
  console.log("Selected!");
};

const SelectBtn = props => {
  return <Button onClick={handleClick}>Select</Button>;
};

export default SelectBtn;
