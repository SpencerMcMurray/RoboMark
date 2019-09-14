import React from "react";
import Display from "../components/Display";
import ScanBtn from "../components/buttons/ScanBtn";

const Pages = () => {
  return (
    <Display
      title="Pages"
      headers={["Name", "Scan"]}
      items={[{ id: 2, name: "Page 1", scan: <ScanBtn id={2} /> }]}
    />
  );
};

export default Pages;
