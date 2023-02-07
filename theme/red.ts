function redHighlightStyle() {
  const config = {
    name: 'red',
    dark: true,
    background: '#390000',
    foreground: '#F8F8F8',
    selection: '#750000',
    cursor: '#970000',
    dropdownBackground: '#390000',
    dropdownBorder: '#220000',
    activeLine: '#ff000033',
    matchingBracket: '#ff000033',
    keyword: '#f12727ff',
    storage: '#ff6262ff',
    variable: '#fb9a4bff',
    parameter: '#fb9a4bff',
    function: '#ffb454ff',
    string: '#cd8d8dff',
    constant: '#994646ff',
    type: '#9df39fff',
    class: '#fec758ff',
    number: '#994646ff',
    comment: '#e7c0c0ff',
    heading: '#fec758ff',
    invalid: '#ffffffff',
    regexp: '#ffb454ff',
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
