import React from "react";
import Brand from "../components/Brand";
import DisplayColumn from "../components/DisplayColumn";
import Layout from "../components/Layout";

const Pages = props => {
  return (
    <Layout>
      <div className="d-flex justify-content-center flex-wrap">
        <div className="row w-100" style={{ height: "90vh" }}>
          <div className="col-lg-3">
            <Brand />
          </div>
          <div className="col-lg-6 d-flex align-items-center">
            <DisplayColumn />
          </div>
          <div className="col-lg-3">
            <Brand />
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default Pages;
