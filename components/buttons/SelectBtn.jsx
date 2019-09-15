import React from "react";
import { Button } from "react-bootstrap";
import Link from "next/link";

const SelectBtn = props => {
  return (
    <Link href={`/questions?test=${props.id}`}>
      <Button variant="success">Questions</Button>
    </Link>
  );
};

export default SelectBtn;
