import React from "react";
import Brand from "../components/Brand";
import DisplayColumn from "../components/DisplayColumn";
import ButtonSpread from "../components/ButtonSpread";
import Layout from "../components/Layout";

import ScanBtn from "../components/buttons/ScanBtn";

const Pages = props => {
  return (
    <Layout>
      <div className="d-flex justify-content-center flex-wrap">
        <div className="row w-100" style={{ height: "90vh" }}>
          <div className="col-lg-3">
            <Brand />
          </div>
          <div className="col-lg-6 d-flex align-items-center">
            <DisplayColumn
              title="Pages"
              headers={["Name", "Scan"]}
              items={[{ id: 2, name: "Page 1", scan: <ScanBtn id={2} /> }]}
            />
          </div>
          <div className="col-lg-3">
            <ButtonSpread />
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default Pages;
