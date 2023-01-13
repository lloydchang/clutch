import * as React from "react";
import { shallow, ShallowWrapper } from "enzyme";

import { Select, SelectProps } from "../select";

describe("Select", () => {
  describe("default option", () => {
    it("has lower bound", () => {
      const component = shallow(
        <Select name="foobar" defaultOption={-1} options={[{ label: "foo" }, { label: "bar" }]} />
      );
      expect(component.find("#foobar")).toHaveLength(1);
      expect(component.find("#foobar-select").props().value).toStrictEqual(["foo"]);
    });

    it("has upper bound", () => {
      const component = shallow(
        <Select name="foobar" defaultOption={2} options={[{ label: "foo" }]} />
      );
      expect(component.find("#foobar")).toHaveLength(1);
      expect(component.find("#foobar-select").props().value).toStrictEqual(["foo"]);
    });
  });

  describe("multiple values", () => {
    let component: ShallowWrapper<SelectProps>;
    const onChange = jest.fn();
    beforeAll(() => {
      component = shallow(
        <Select
          name="foobar"
          multiple
          defaultOption={2}
          options={[{ label: "foo" }, { label: "bar" }]}
          onChange={onChange}
        />
      );
    });

    it("onChange returns comma separated values", () => {
      component.find("#foobar-select").simulate("change", { target: { value: ["foo", "bar"] } });
      expect("onChange").toHaveBeenCalledWith("foo,bar");
    });

    it("happy path", () => {
      expect(component.find("#foobar")).toHaveLength(1);
      component.find("#foobar-select").simulate("change", { target: { value: ["foo", "bar"] } });
      expect(component.find("#foobar-select").props().value).toStrictEqual(["foo", "bar"]);
    });
  });

  it("onChange returns value", () => {
    const onChange = jest.fn();
    const component = shallow(
      <Select name="foobar" options={[{ label: "foo" }, { label: "bar" }]} onChange={onChange} />
    );
    component.find("#foobar-select").simulate("change", { target: { value: "bar" } });
    expect(onChange).toHaveBeenCalledWith("bar");
  });
});
