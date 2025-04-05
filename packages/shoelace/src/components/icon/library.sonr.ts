import { getBasePath } from '../../utilities/base-path.js';
import type { IconLibrary } from './library.js';

const library: IconLibrary = {
  name: 'sonr',
  resolver: name => getBasePath(`assets/sonr/${name}.svg`)
};

export default library;
