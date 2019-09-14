import React from "react";
import Display from "../components/Display";
import ScanBtn from "../components/buttons/ScanBtn";
import { Button } from "react-bootstrap";

const btns = [
  <Button
    variant="success"
    onClick={() => {
      console.log("Created new");
    }}
  >
    Create New
  </Button>,
  <Button
    onClick={() => {
      console.log("Finshed");
    }}
  >
    Finish
  </Button>
];

const Pages = () => {
  return (
    <Display
      title="Pages"
      headers={["Number", "Scan"]}
      items={[{ id: 2, name: 1, scan: <ScanBtn id={2} /> }]}
      buttons={btns}
    />
  );
};

export default Pages;
