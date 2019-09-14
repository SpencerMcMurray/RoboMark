import React from "react";
import { Button } from "react-bootstrap";

const handleClick = () => {
  console.log("Scaning!");
};

const ScanBtn = props => {
  return (
    <Button variant="info" onClick={handleClick}>
      Scan
    </Button>
  );
};

export default ScanBtn;
