function lightPlusHighlightStyle() {
  const config = {
    name: 'lightPlus',
    dark: false,
    background: '#ffffff',
    foreground: '#333333',
    selection: '#ADD6FF',
    cursor: '#333333',
    dropdownBackground: '#F3F3F3',
    dropdownBorder: '#C8C8C8',
    activeLine: '#ADD6FF',
    matchingBracket: '#0064001a',
    keyword: '#0000ff',
    storage: '#0000ff',
    variable: '#001080',
    parameter: '#001080',
    function: '#795E26',
    string: '#a31515',
    constant: '#0000ff',
    type: '#267f99',
    class: '#267f99',
    number: '#098658',
    comment: '#008000',
    heading: '#000080',
    invalid: '#cd3131',
    regexp: '#811f3f',
  };
  return HighlightStyle.define([ 
    {tag: t.keyword, color: config.keyword},
    {tag: [t.name, t.deleted, t.character, t.macroName], color: config.variable},
    {tag: [t.propertyName], color: config.function},
    {tag: [t.processingInstruction, t.string, t.inserted, t.special(t.string)], color: config.string},
    {tag: [t.function(t.variableName), t.labelName], color: config.function},
    {tag: [t.color, t.constant(t.name), t.standard(t.name)], color: config.constant},
    {tag: [t.definition(t.name), t.separator], color: config.variable},
    {tag: [t.className], color: config.class},
    {tag: [t.number, t.changed, t.annotation, t.modifier, t.self, t.namespace], color: config.number},
    {tag: [t.typeName], color: config.type, fontStyle: config.type},
    {tag: [t.operator, t.operatorKeyword], color: config.keyword},
    {tag: [t.url, t.escape, t.regexp, t.link], color: config.regexp},
    {tag: [t.meta, t.comment], color: config.comment},
    {tag: t.strong, fontWeight: 'bold'},
    {tag: t.emphasis, fontStyle: 'italic'},
    {tag: t.link, textDecoration: 'underline'},
    {tag: t.heading, fontWeight: 'bold', color: config.heading},
    {tag: [t.atom, t.bool, t.special(t.variableName)], color: config.variable},
    {tag: t.invalid, color: config.invalid},
    {tag: t.strikethrough, textDecoration: 'line-through'},
  ]);
}
