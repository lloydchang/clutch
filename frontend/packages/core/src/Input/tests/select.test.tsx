import * as React from "react";
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";

import "@testing-library/jest-dom";

import { Select } from "..";

test("has lower bound", () => {
  const { container } = render(
    <Select name="foobar" defaultOption={-1} options={[{ label: "foo" }, { label: "bar" }]} />
  );

  expect(container.querySelector("#foobar-select")).toBeInTheDocument();
  expect(container.querySelector("#foobar-select")).toHaveTextContent("foo");
});

test("has upper bound", () => {
  const { container } = render(
    <Select name="foobar" defaultOption={2} options={[{ label: "foo" }]} />
  );

  expect(container.querySelector("#foobar-select")).toBeInTheDocument();
  expect(container.querySelector("#foobar-select")).toHaveTextContent("foo");
});

test("onChange returns single values", async () => {
  const onChange = jest.fn();

  const { getByRole } = render(
    <Select
      name="foobar"
      defaultOption={0}
      options={[{ label: "foo" }, { label: "bar" }]}
      onChange={onChange}
    />
  );

  const user = userEvent.setup();

  await user.click(getByRole("button"));

  // select bar
  await user.selectOptions(screen.getByRole("listbox"), [
    screen.getByRole("option", { name: "bar" }),
  ]);

  expect(onChange).toHaveBeenCalledTimes(2);
  expect(onChange.mock.calls[onChange.mock.calls.length - 1][0]).toBe("bar");
});

test("onChange returns multiple values", async () => {
  const onChange = jest.fn();

  const { getByRole } = render(
    <Select
      name="foobar"
      multiple
      defaultOption={0}
      options={[{ label: "foo" }, { label: "bar" }]}
      onChange={onChange}
    />
  );

  const user = userEvent.setup();

  await user.click(getByRole("button"));
  // deselect default
  await user.selectOptions(screen.getByRole("listbox"), [
    screen.getByRole("option", { name: "foo" }),
  ]);
  // select both options
  await user.selectOptions(screen.getByRole("listbox"), [
    screen.getByRole("option", { name: "foo" }),
    screen.getByRole("option", { name: "bar" }),
  ]);

  expect(onChange).toHaveBeenCalledTimes(3);
  expect(onChange.mock.calls[onChange.mock.calls.length - 1][0]).toBe("foo,bar");
});
