import React from "react";
import Display from "../components/Display";
import ScanBtn from "../components/buttons/ScanBtn";
import PageBtns from "../components/buttons/PageBtns";

const Pages = () => {
  return (
    <Display
      title="Pages"
      headers={["Number", "Scan"]}
      items={[{ id: 2, name: 1, scan: <ScanBtn id={2} /> }]}
      buttons={<PageBtns />}
    />
  );
};

export default Pages;
